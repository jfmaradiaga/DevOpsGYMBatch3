package main

import (
    "fmt"
    "net/http"
    "os"
    "html/template"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

type HelloRequest struct {
    Name string `json:"name"`
}

var createdInstanceID string // Declare this at the global scope

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    if err := r.ParseForm(); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, "Invalid form data")
        return
    }

    name := r.FormValue("name")

    if name == "" {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, "Name parameter is missing")
        return
    }

    response := fmt.Sprintf("Hello, %s!", name)
    fmt.Fprintln(w, response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    tmpl := template.Must(template.ParseFiles("index.html"))
    if err := tmpl.Execute(w, nil); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintln(w, "Internal server error")
        return
    }
}

func createEC2Instance(sess *session.Session) {
    svc := ec2.New(sess)
    input := &ec2.RunInstancesInput{
        ImageId:      aws.String("ami-053b0d53c279acc90"),
        InstanceType: aws.String("t2.micro"),
        // Your EC2 configuration here, like AMI ID, instance type, etc.
        MinCount: aws.Int64(1),
        MaxCount: aws.Int64(1),
    }

    result, err := svc.RunInstances(input)
    if err != nil {
        fmt.Println("Error creating instance:", err)
        return
    }

    createdInstanceID = *result.Instances[0].InstanceId
}

func terminateEC2Instance(sess *session.Session) {
    svc := ec2.New(sess)
    input := &ec2.TerminateInstancesInput{
        InstanceIds: []*string{
            aws.String(createdInstanceID),
        },
    }

    _, err := svc.TerminateInstances(input)
    if err != nil {
        fmt.Println("Error terminating instance:", err)
    }
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    createEC2Instance(sess)
    fmt.Fprintln(w, "EC2 Instance Created with ID:", createdInstanceID)
}

func terminateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    terminateEC2Instance(sess)
    fmt.Fprintln(w, "EC2 Instance Terminated with ID:", createdInstanceID)
}

var sess *session.Session

func main() {
    var err error

    // Initialize AWS Session
    sess, err = session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
        Credentials: credentials.NewStaticCredentials("AKIAXQHXIJL2IMWQMXXY", "ckYp4OaWcuSi/3ach5NvivgaDPIqn20JybqCeBUV", ""),
    })
    if err != nil {
        fmt.Println("Error initializing AWS session:", err)
        return
    }

    if len(os.Args) < 2 {
        fmt.Println("Usage: ./cli <host:port>")
        return
    }

    host := os.Args[1]

    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/api/hello", helloHandler)
    http.HandleFunc("/api/ec2/create", createHandler)
    http.HandleFunc("/api/ec2/terminate", terminateHandler)

    err = http.ListenAndServe(host, nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

