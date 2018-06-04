CREATE TABLE IF NOT EXISTS dogs
(
id SERIAL,
name TEXT NOT NULL,
energy_level INTEGER NOT NULL,
CONSTRAINT dogs_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS monkeys
(
id SERIAL,
name TEXT NOT NULL,
energy_level INTEGER NOT NULL,
CONSTRAINT dogs_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS weapons
(
id SERIAL,
name TEXT NOT NULL,
energy_level INTEGER NOT NULL,
CONSTRAINT weapons_pkey PRIMARY KEY (id)
);

INSERT INTO dogs(name, energy_level) VALUES('Test Dog 1', 1);
INSERT INTO dogs(name, energy_level) VALUES('Test Dog 2', 2);
INSERT INTO dogs(name, energy_level) VALUES('Test Dog 3', 3);
INSERT INTO dogs(name, energy_level) VALUES('Test Dog 4', 4);
INSERT INTO dogs(name, energy_level) VALUES('Test Dog 5', 5);



INSERT INTO monkeys(name, energy_level) VALUES('Test Monkey 1', 1);
INSERT INTO monkeys(name, energy_level) VALUES('Test Monkey 2', 2);
INSERT INTO monkeys(name, energy_level) VALUES('Test Monkey 3', 3);
INSERT INTO monkeys(name, energy_level) VALUES('Test Monkey 4', 4);
INSERT INTO monkeys(name, energy_level) VALUES('Test Monkey 5', 5);

INSERT INTO weapons(name, power_level) VALUES('Test Weapon 1', 1);
INSERT INTO weapons(name, power_level) VALUES('Test Weapon 2', 2);
INSERT INTO weapons(name, power_level) VALUES('Test Weapon 3', 3);
INSERT INTO weapons(name, power_level) VALUES('Test Weapon 4', 4);
INSERT INTO weapons(name, power_level) VALUES('Test Weapon 5', 5);

