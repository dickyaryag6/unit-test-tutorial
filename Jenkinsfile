pipeline {
  agent any
  
  stages {
    stage ('Run Unit Test') {
        steps {
          sh 'go test ./... -v'
        }
    }
    stage ('Build') {
        steps {
          sh 'go build'
        }
    }
  }
}
