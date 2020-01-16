# gql-tuts
tutorial source:
https://dev.to/cmelgarejo/creating-an-opinionated-graphql-server-with-go-part-1-3g3l

make build script executable if not already. Example below </br>
`$ chmod +x scripts/build.sh`
after that build server <br/>
`$ ./scripts/build.sh` <br/>
if successful, run server <br>
`$ ./build/gql-server`

To run without building run: </br>
`$ ./scripts/run.sh`
make sure run.sh is executable
I have commented out env variable export and unset because I am using my local .env file.