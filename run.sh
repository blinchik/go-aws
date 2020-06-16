SSH_LOCAL="${HOME}/.ssh/"
SSH_DOCKER="${whoami}/.ssh/"


docker run -i \
    -e aws_region=$aws_region \
    -e user=${whoami} \
    -v "$SSH_LOCAL:$SSH_DOCKER" \
    -u `stat -c "%u:%g" $SSH_LOCAL` \
    devblinchik/go-aws:latest $1 $2
