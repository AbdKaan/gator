# Gator

Requirements:
- Go
- PostgreSQL

Make sure to run postgres on port 5432.

## Installation:

Clone the repository or download ZIP file and extract them.
```
git clone https://github.com/AbdKaan/gator.git
```

Go into the package's folder and install using go.

```
go install
```

You need to create a file named `.gatorconfig.json` on your home directory.
For windows this would be `C:\Users\<User-Name>\`, for linux: `~/`

It needs to include:

```
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable")
```

## Usage:

Once you are done with installation, you can start by registering with a user name.

```
gator register <user_name>
```

Now you can add feeds you want.

```
gator addfeed <name> <url>
```

Example:

```
gator addfeed "TechCrunch" "https://techcrunch.com/feed/"
```

After adding the feeds you want, you can start aggregating posts from them by using the command `agg` and specifying the time_interval of fetching the titles.

Example for time_interval parameter: `1h2m3s`

This will keep fetching posts from the latest feeds by order.

```
gator agg <time_interval>
```

Once you run the above command, you need to open another terminal to browse them.

Now you can browse the posts you have fetched using `browse` command. It has an optional `limit` parameter (default is 2) that specifies the number of posts you will get. It will fetch the latest updated `limit` amount of posts.

```
gator browse <limit>
```

### Commands

- `register <user_name>`: register with a user name of your choice, it will automatically login after registering
- `login <user_name>`: login to another user account
- `users`: display users
- `addfeed <name> <url>`: add a feed by providing it a name and it's url
- `feeds`: display added feeds
- `follow <url>`: follow a feed using it's url. this command is useful to follow a feed that another user added or to follow a feed that you unfollowed before
- `following`: display the feeds you are following
- `unfollow <url>`: unfollow the given feed url
- `agg <time_interval>`: aggregate the posts from the feeds user follows
- `browse <limit>`: browse the aggregated posts sorted by the latest ones
- `reset`: deletes every data