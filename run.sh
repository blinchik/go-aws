SSH_LOCAL="${HOME}/.ssh/"
SSH_DOCKER="/home/root/.ssh/"


docker run -i \
    -e aws_region=$aws_region \
    -v "/$SSH_LOCAL:$SSH_DOCKER" \
    devblinchik/go-aws:latest $1 $2
