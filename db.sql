CREATE TABLE `person` (
`id` int(6) unsigned NOT NULL AUTO_INCREMENT,
`identity` varchar(30) NOT NULL,
`fullname` varchar(30) NOT NULL,
`address` varchar(50) NOT NULL,
`telnum`varchar(14) NOT NULL,
`email` varchar(30) NOT NULL, PRIMARY KEY (`id`) )
ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1