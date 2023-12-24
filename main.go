pipeline {
  agent any
  stages {
    stage('Test') {
      steps {
        sh 'go test .'
     }
    }
    stage('Build') {
      steps {
        sh 'docker build . -t ubuntu-bionic:8082/hello-world:v$BUILD_NUMBER'
      }
    }
  }
}
