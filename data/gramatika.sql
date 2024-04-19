BEGIN;

CREATE TABLE nor (
    nor TEXT PRIMARY KEY
);
Insert into nor (nor) values ('Ni');
Insert into nor (nor) values ('Hura');
Insert into nor (nor) values ('Gu');
Insert into nor (nor) values ('Zu');
Insert into nor (nor) values ('Zuek');
Insert into nor (nor) values ('Haiek');

CREATE TABLE nori (
    nori TEXT PRIMARY KEY
);
Insert into nori (nori) values ('Niri');
Insert into nori (nori) values ('Hari');
Insert into nori (nori) values ('Guri');
Insert into nori (nori) values ('Zuri');
Insert into nori (nori) values ('Zuei');
Insert into nori (nori) values ('Haiei');

CREATE TABLE nork (
    nork TEXT PRIMARY KEY
);
Insert into nork (nork) values ('Nik');
Insert into nork (nork) values ('Hark');
Insert into nork (nork) values ('Guk');
Insert into nork (nork) values ('Zuk');
Insert into nork (nork) values ('Zuek');
Insert into nork (nork) values ('Haiek');

CREATE TABLE denbora (
    denbora TEXT PRIMARY KEY
);
Insert into denbora (denbora) values ('Orainaldia');
Insert into denbora (denbora) values ('Lehenaldia');

Create table nor_nork (
    nor TEXT,
    nork TEXT,
    denbora TEXT,
    aditz_languzailea TEXT,
    PRIMARY KEY (nor, nork, denbora),
    FOREIGN KEY (nor) REFERENCES nor(nor),
    FOREIGN KEY (nork) REFERENCES nork(nork),
    FOREIGN KEY (denbora) REFERENCES denbora(denbora)
);

Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Ni', 'Hark', 'Nau', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Ni', 'Zuk', 'Nauzu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Ni', 'Zuek', 'Nauzue', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Ni', 'Haiek', 'Naute', 'Orainaldia');


Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Hura', 'Nik', 'Dut', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Hura', 'Hark', 'Du', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Hura', 'Guk', 'Dugu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Hura', 'Zuk', 'Duzu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Hura', 'Zuek', 'Duzue', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_languzailea, denbora) values ('Hura', 'Haiek', 'Dute', 'Orainaldia');








Create table nor_nori (
    nor TEXT,
    nori TEXT,
    denbora TEXT,
    aditz_languzailea TEXT,
    PRIMARY KEY (nor, nori, denbora),
    FOREIGN KEY (nor) REFERENCES nor(nor),
    FOREIGN KEY (nori) REFERENCES nori(nori),
    FOREIGN KEY (denbora) REFERENCES denbora(denbora)
);

Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Niri', 'Zait', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Hari', 'Zaio', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Guri', 'Zaigu', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Zuri', 'Zaizu', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Zuei', 'Zaizue', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Haiei', 'Zaie', 'Orainaldia');

Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Niri', 'Zaizkit', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Hari', 'Zaizkio', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Guri', 'Zaizkigu', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Zuri', 'Zaizkizu', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Zuei', 'Zaizkizue', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Haiei', 'Zaizkie', 'Orainaldia');

Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Niri', 'Zitzaidan', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Hari', 'Zitzaion', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Guri', 'Zitzaigun', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Zuri', 'Zitzaizun', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Zuei', 'Zitzaizuen', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Hura', 'Haiei', 'Zitzaien', 'Lehenaldia');

Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Niri', 'Zitzazkidan', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Hari', 'Zitzaizkion', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Guri', 'Zitzaizkigun', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Zuri', 'Zitzaizkizun', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Zuei', 'Zitzaizkizuen', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_languzailea, denbora) values ('Haiek', 'Haiei', 'Zitzaizkien', 'Lehenaldia');


Create table nor_nori_nork (
    nor TEXT,
    nori TEXT,
    nork TEXT,
    denbora TEXT,
    aditz_languzailea TEXT,
    PRIMARY KEY (nor, nori, nork, denbora),
    FOREIGN KEY (nor) REFERENCES nor(nor),
    FOREIGN KEY (nori) REFERENCES nori(nori),
    FOREIGN KEY (nork) REFERENCES nork(nork),
    FOREIGN KEY (denbora) REFERENCES denbora(denbora)
);

-- Insert into nor_nori_nork (nor, nori, nork, aditz_languzailea, denbora) values ('Hura', 'Niri', 'Nik', 'Zait', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_languzailea, denbora) values ('Hura', 'Hari', 'Nik', 'Diot', 'Orainaldia');
-- Insert into nor_nori_nork (nor, nori, nork, aditz_languzailea, denbora) values ('Hura', 'Guri', 'Nik', 'Zaigu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_languzailea, denbora) values ('Hura', 'Zuri', 'Nik', 'Dizut', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_languzailea, denbora) values ('Hura', 'Zuei', 'Nik', 'Dizuet', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_languzailea, denbora) values ('Hura', 'Haiei', 'Nik', 'Diet', 'Orainaldia');


Create view aditz_lagunak(forme, nor, nori, nork, aditz_laguntzilea, denbora) as 
select 'nor_nork', nor, '', nork, aditz_languzailea, denbora from nor_nork
union
select 'nor_nori', nor, nori, '', aditz_languzailea, denbora from nor_nori
union
select 'nor_nori_nork', nor, nori, nork, aditz_languzailea, denbora from nor_nori_nork;


COMMIT;
