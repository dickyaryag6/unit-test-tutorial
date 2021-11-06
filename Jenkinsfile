pipeline {
  agent any
  
  stages {
    stage ('Run Unit Test') {
        steps {
          try {
            sh 'go test ./... -v'
          } catch (e) {
            currentBuild.result = 'ABORTED'
            error("Aborting the build.")
          }
        }
    }
    stage ('Build') {
        steps {
          try {
            sh 'go build'
          } catch (e) {
            currentBuild.result = 'ABORTED'
            error("Aborting the build.")
          }
        }
    }
  }
}
