pipeline {
    agent { label 'dockerserver' }
    stages { 
        stage('build') {
            agent {
                docker {
                // Set both label and image
                    label 'dockerserver'
                    image 'golang'
                 }
            }
            steps {
                sh 'go version'
            }
        }
    }
}
