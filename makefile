ACCOUNT = <your-google-account>
PROJECT = <your-gcp-project>
VERSION = <your-gcp-project-version>

set_config:
	gcloud config set account $(ACCOUNT)
	gcloud config set project $(PROJECT)

run:
	dev_appserver.py app.yaml --skip_sdk_update_check=yes --host 0.0.0.0 --enable_sendmail=yes

update_frontend:
	goapp deploy -application $(PROJECT) -version $(VERSION) app.yaml

update: update_frontend
