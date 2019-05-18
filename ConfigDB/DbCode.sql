create database plpInfo;

use plpInfo;

create table NamesMan (
	IdM int NOT NULL AUTO_INCREMENT,
	FirstName varchar(300) NOT NULL,
    LastName varchar(300) NOT NULL,
    primary key(IdM)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;

create table NamesWoman (
	IdW int NOT NULL AUTO_INCREMENT,
	FirstName varchar(300) NOT NULL,
    LastName varchar(300) NOT NULL,
    primary key(IdW)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;


