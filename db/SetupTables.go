package db

import "log"

func CreateTables() {
	createTables := `
    CREATE TABLE IF NOT EXISTS "public"."user" (
        id SERIAL PRIMARY KEY,
        username VARCHAR NOT NULL UNIQUE,
        password VARCHAR NOT NULL,
        title VARCHAR NOT NULL,
        full_name VARCHAR NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS "public"."token" (
        id SERIAL PRIMARY KEY,
        refresh_token VARCHAR NOT NULL,
        valid_from TIMESTAMP WITH TIME ZONE NOT NULL,
        valid_to TIMESTAMP WITH TIME ZONE NOT NULL,
        user_id INTEGER NOT NULL,
        enabled BOOLEAN NOT NULL DEFAULT TRUE,
        CONSTRAINT fk_user FOREIGN KEY (user_id)
            REFERENCES "user" (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
    );

    CREATE TABLE IF NOT EXISTS "public"."blogpost" (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        text TEXT NOT NULL,
        title TEXT NOT NULL,
        author TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_user FOREIGN KEY (user_id)
            REFERENCES "user" (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
    );
        CREATE TABLE IF NOT EXISTS "public"."blogpost_tags" (
        id SERIAL PRIMARY KEY,
        tag TEXT NOT NULL,
        blogpost_id INTEGER NOT NULL,
        CONSTRAINT fk_blog FOREIGN KEY (blogpost_id)
            REFERENCES "blogpost" (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
    );
    
    `

	if _, err := DB.Exec(createTables); err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}
}
