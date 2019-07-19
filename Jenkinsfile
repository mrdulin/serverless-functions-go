// May be the robustest Kubernetes deployment tool here.
// Indent standard make sure you are following 'martinda/Jenkinsfile-vim-syntax' 'modille/groovy.vim' before commit changes.
pipeline {
    agent any
    parameters {
        choice(name: "GCP_ENV",
        choices: ["dev", "qa", "stage", "prod"],
        description: "pg-gx-e-app-700458 - DEV , pg-us-n-app-991483 - QA, pg-us-p-app-319619 - STAGE pg-us-p-app-695770 - PROD")

        string(name: "ZONE",
        defaultValue: "us-east1-c",
        description: "Deploy any region/zone you want")

        string(name: "BRANCH",
        defaultValue: "develop",
        description: "Deploy any branch you want")

        string(name: "COMMIT",
        defaultValue: "",
        description: "Leave it blank if no specific commit or TAG you want to checkout")
    }

    /* Project level isolation */
    environment {
        GIT_REPO = "serverless-functions-go"
        ORGANIZATION = "NV-Cedar"
        PROJECT_NAME = "cedar"
        DOCKER_REGISTRY= "dtr.artifacts.pwc.com/pwcusnv"
    }

    stages {

        stage("Log parameters"){
            steps {
                echo "${GCP_ENV} ${ZONE} ${GIT_REPO} BRANCH=${BRANCH} COMMIT=${COMMIT}"
            }
        }

        stage("Checkout"){
            steps {
                cleanWs()
                git branch: "${BRANCH}",
                credentialsId: "jenkins-gitlab-ssh",
                url: "git@github.pwc.com:${ORGANIZATION}/${GIT_REPO}"

                sh "git checkout ${COMMIT}"
                sh "find . -maxdepth 2"

                script {
                    sh "mkdir -p ${JENKINS_HOME}/userContent/${GIT_REPO}"
                    def COMMIT7 = sh( script: "git rev-parse HEAD | cut -c 1-7 2>&1 | tee ${JENKINS_HOME}/userContent/${GIT_REPO}/COMMIT.txt", returnStdout: true)
                    /* COMMIT7 == ${COMMIT:1:7} */
                    echo "COMMIT7 = ${COMMIT7}"
                    sh "env"
                }
            }
        }

        stage("Deploy"){
            steps {
                script {
                    def COMMIT7 = sh ( script: "cat ${JENKINS_HOME}/userContent/${GIT_REPO}/COMMIT.txt", returnStdout: true).trim()

                    if ( GCP_ENV == "dev" ){
                        GCP_PROJECT_ID = "pg-gx-e-app-700458"
                        instance_name = "pg-gx-e-app-700458:us-central1:cedar-database"
                    }else if ( GCP_ENV == "qa" ){
                        GCP_PROJECT_ID = "pg-us-n-app-991483"
                        instance_name = "pg-us-n-app-991483:us-central1:cedar-database"
                    }else if ( GCP_ENV == "stage" ){
                        GCP_PROJECT_ID = "pg-us-p-app-319619"
                        instance_name = "pg-us-p-app-319619:us-east1:cedar-database"
                    }else if ( GCP_ENV == "prod" ){
                        GCP_PROJECT_ID = "pg-us-p-app-695770"
                        instance_name = "pg-us-p-app-695770:us-east1:cedar-database"
                    }
                    withCredentials([file(credentialsId: "${GCP_PROJECT_ID}-jenkins-keys", variable: "SERVICE_ACCOUNT")]) {
                        sh "gcloud auth activate-service-account --key-file ${SERVICE_ACCOUNT}"
                        sh "gcloud config set project ${GCP_PROJECT_ID}"
                        sh "gcloud config set compute/zone ${ZONE}"
                        sh "gcloud beta run deploy --image=${DOCKER_REGISTRY}/${PROJECT_NAME}/${GIT_REPO}:${COMMIT7} --set-cloudsql-instances ${instance_name} --set-env-vars ENV=production,GCP_PROJECT=${GCP_PROJECT_ID} --platform managed --region ${ZONE}"
                    }
                }
            }
        }
    }
}

