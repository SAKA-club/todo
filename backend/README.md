#Deployment
We are using [fly.io](https://fly.io/docs/speedrun/) for deployment

In order to deploy to fly.io flyctl needs to be installed

Mac:
``brew install flyctl``

Windows:
`iwr https://fly.io/install.ps1 -useb | iex`

## flyctl commands

in order to deploy
`flyctl launch`

Follow commands and keep track of the postgres credentials, they will be required to create a db_url in order to connect to the database using dbmate

Once finished a fly.toml file will be created 