pipeline {
    agent { label 'dockerserver' }
    stages {
        agent {
        docker {
          // Set both label and image
          label 'dockerserver'
          image 'golang'
        }
      }
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}
