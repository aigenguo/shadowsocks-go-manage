CREATE TABLE IF NOT EXISTS members (
		id int(10) unsigned NOT NULL AUTO_INCREMENT,
		port int(11) NOT NULL,
		password varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		name varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		created_at timestamp NULL DEFAULT NULL,
		updated_at timestamp NULL DEFAULT NULL,
		PRIMARY KEY (id),
	  	UNIQUE KEY members_port_unique (port)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS migrations (
		migration varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		batch int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS users (
		id int(10) unsigned NOT NULL AUTO_INCREMENT,
		email varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		password varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		remember_token varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
		created_at timestamp NULL DEFAULT NULL,
		updated_at timestamp NULL DEFAULT NULL,
		PRIMARY KEY (id),
		UNIQUE KEY users_email_unique (email)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;