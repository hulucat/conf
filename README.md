# conf
Loading config file from running dir.
The config file should named .env.conf, formated as: 

############### example start ###############
mysql = user1:password1@tcp(127.0.0.1:3306)/db1?autocommit=true <br/>
log_path = /var/log/abc.log
############### example end   ###############

So the conf will load these config items into key-value pair map, and 
you will get them through conf.Get("mysql"), conf.Get("log_path").