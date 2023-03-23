FROM mysql:5.7.22

COPY /deployments/custom.cnf /etc/mysql/conf.d/custom.cnf