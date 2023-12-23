//package main

//import (
//	"fmt"
//	"log"
//	"os"
//)

//func hello(name string) string {
//	return fmt.Sprintf("Hello %s!", name)
//}

//func main() {
//	if len(os.Args) < 2 {
//		log.Fatalln("Missing name: hello <name>")
//	}
//	fmt.Println(hello(os.Args[1]))
//}

pipeline {
 agent any
 stages {
  stage('Git') {
   steps {git 'https://github.com/Evgenii-379/git-homework-8-02.md.git'}
  }
  stage('Test') {
   steps {
    sh '/usr/local/go/bin/go test .'
   }
  }
  stage('Build') {
   steps {
    sh 'docker build . -t ubuntu-bionic:8082/hello-world:v$BUILD_NUMBER'
   }
  }
  stage('Push') {
   steps {
    sh 'docker login ubuntu-bionic:8082 -u admin -p admin && docker push ubuntu-bionic:8082/hello-world:v$BUILD_NUMBER && docker logout'   }
  }
 }
}

