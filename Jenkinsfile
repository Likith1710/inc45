pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIALS = 'docker_creds'
    }

    stages {
        stage('GO changes') {
            when {
                changeset "**/go/**"
            }
            steps {
                sh 'echo Change in GO'
                dir('go'){
                    sh 'ls'
                    sh 'docker build -t goapplication .'
                    sh 'docker tag goapplication likith1710/inc42test'
                    withCredentials([usernamePassword(credentialsId: DOCKER_HUB_CREDENTIALS, usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        sh 'docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD'
                        sh 'docker push likith1710/inc42test'
                    }
                }
            }
        }
        stage('Next.js changes') {
            when {
                changeset "**/nextjs/**"
            }
            steps {
                sh 'echo Change in Next.js'
                dir('nextjs') {
                    sh 'ls'
                    sh 'docker build -t nextjsapplication .'
                    sh 'docker tag nextjsapplication likith1710/nexjsapplication'
                    sh 'docker push likith1710/nexjsapplication'
                }
            }
        }
        stage('Wordpress changes') {
            when {
                changeset  "**/wordpress/**"
            }
            steps {
                sh 'echo Change in Wordpress'
                dir('wordpress') {
                    sh 'docker build -t wordpresssite .'
                    sh 'docker tag wordpresssite likith1710/wordpress'
                    sh 'docker push likith1710/wordpress'
                }
            }
        }
        stage('Compose Build') {
            steps {
                sh'docker-compose down'
                sh 'docker-compose up -d'
            }
        }
    }
}

