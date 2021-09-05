Command line tool to check the validity of Jet Golang Templates
==================================================================

Checks the validity of jet templates and prints the errors if found in stderr.

Could me useful to add it to your build pipelines to make sure that all your templates are valid.

#### Usage

Assuming that your templates are in folder `/home/user/templates`

##### Option 1 (Use Dockerhub image)

```
docker run -v /home/user/templates:/templates gosom/check-golang-templates check-golang-templates -folder /templates
```

##### Option 2 (Download binary [ Linux Only ])

Download the latest binary from the releases: https://github.com/gosom/check-golang-templates/releases/tag/v1.0.0

Make it executable if needed and run

`check-golang-templates -folder /home/user/templates`

##### Option 3(Build on your own)

```
git clone git@github.com:gosom/check-golang-templates.git
cd check-golang-template
go mod download
go install
```

`check-golang-templates -folder /home/user/templates`
