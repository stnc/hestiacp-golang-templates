git add .
git commit -m "feat: auto"
git push

rm -rf main

# GOOS=linux GOARCH=386 go build main.go
GOOS=linux GOARCH=amd64 go build main.go

tar -cvzf view.tar.gz public/locales/* public/view/* main

mkdir CI

rm -rf CI/*

mv view.tar.gz CI/view.tar.gz

scp -i oracle.key view.tar.gz ubuntu@141.144.242.109:/home/ubuntu