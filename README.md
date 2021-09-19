## This for base golang command project.


### Build
```bash
git clone https://github.com/neilkuan/base-go-cmd.git

go build 
```

### Get
```bash
cp sample-v.json videos.json
./base-go-cmd get --all
```

### Save
```bash
./base-go-cmd add -id "001" -title "apple tv" -url "https://youtu.be/apple" -imageurl https://google.com -desc "apple"
```

### Get One
```bash
./base-go-cmd get --id 001
```