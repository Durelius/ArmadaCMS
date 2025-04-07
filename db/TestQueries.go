package db

import "log"

func TestQuery() string {
	var test string
	testSelectSQL := "SELECT test FROM testtable WHERE id = $1"
	err := DB.Get(&test, testSelectSQL, 1)
	if err != nil {
		log.Fatal(err)
	}
	return test

}

func CreateTestTable() {
	createTestTableSQL := `CREATE SEQUENCE IF NOT EXISTS untitled_table_209_id_seq;

CREATE TABLE "public"."testtable" (
    "id" int4 NOT NULL DEFAULT nextval('untitled_table_209_id_seq'::regclass),
    "test" varchar NOT NULL,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX untitled_table_209_pkey ON public.testtable USING btree (id);

INSERT INTO "public"."testtable" ("id", "test") VALUES
(1, 'Successful content delivery');
`
	DB.Exec(createTestTableSQL)

}
