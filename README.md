## Without docker-compose on linux OS:
Requirement :
1. reflex : go get github.com/cespare/reflex
2. install these on linux : (autoconf automake libtool gettext gettext-dev make g++)
3. run : make watch

## With docker-compose :
1. run : docker-compose up -d app-development
2. tailing logs : docker-compose logs --tail=100 -f app-development
