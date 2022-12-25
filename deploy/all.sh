git add .
git commit -m "feat: auto"
git push

rm -rf main

# GOOS=linux GOARCH=386 go build main.go
GOOS=linux GOARCH=amd64 go build main.go

tar -cvzf all.tar.gz public/static/* public/locales/* public/view/*  other/hestiaGO/* main

mkdir CI

rm -rf CI/*

mv all.tar.gz CI/all.tar.gz


scp -i oracle.key all.tar.gz ubuntu@141.144.242.109:/home/ubuntu