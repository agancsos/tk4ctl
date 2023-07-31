from debian:11
run apt-get update -y
run apt-get install -y python3 python3-pip libcurl4-openssl-dev golang openssh-server cmake vim
run python3 -m pip install requests boto3
run mkdir -p /root/stuff/scripts/go/tk4ctl
copy . /root/stuff/scripts/go/tk4ctl
entrypoint ["bash", "-c", "cd /root/stuff/scripts/go/tk4ctl; make; tail -f /dev/null"]
