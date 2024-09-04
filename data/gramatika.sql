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
    aditz_lagunzailea TEXT,
    PRIMARY KEY (nor, nork, denbora),
    FOREIGN KEY (nor) REFERENCES nor(nor),
    FOREIGN KEY (nork) REFERENCES nork(nork),
    FOREIGN KEY (denbora) REFERENCES denbora(denbora)
);

Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Ni', 'Hark', 'Nau', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Ni', 'Zuk', 'Nauzu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Ni', 'Zuek', 'Nauzue', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Ni', 'Haiek', 'Naute', 'Orainaldia');

Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Hura', 'Nik', 'Dut', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Hura', 'Hark', 'Du', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Hura', 'Guk', 'Dugu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuk', 'Duzu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuek', 'Duzue', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Hura', 'Haiek', 'Dute', 'Orainaldia');

Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Gu', 'Hark', 'Gaitu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Gu', 'Zuk', 'Gaituzu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Gu', 'Zuek', 'Gaituzue', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Gu', 'Haiek', 'Gaituzte', 'Orainaldia');

Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Zu', 'Nik', 'Zaitut', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Zu', 'Hark', 'Zaitu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Zu', 'Guk', 'Zaitugu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Zu', 'Haiek', 'Zaituzte', 'Orainaldia');

Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Zuek', 'Nik', 'Zaituztet', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Zuek', 'Hark', 'Zaituzte', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Zuek', 'Guk', 'Zaituztegu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Zuek', 'Haiek', 'Zaituztete', 'Orainaldia');

Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Nik', 'Ditut', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Hark', 'Ditu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Guk', 'Ditugu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuk', 'Dituzu', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuek', 'Dituzue', 'Orainaldia');
Insert into nor_nork (nor, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Haiek', 'Dituzte', 'Orainaldia');








Create table nor_nori (
    nor TEXT,
    nori TEXT,
    denbora TEXT,
    aditz_lagunzailea TEXT,
    PRIMARY KEY (nor, nori, denbora),
    FOREIGN KEY (nor) REFERENCES nor(nor),
    FOREIGN KEY (nori) REFERENCES nori(nori),
    FOREIGN KEY (denbora) REFERENCES denbora(denbora)
);

Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Niri', 'Zait', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Hari', 'Zaio', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Guri', 'Zaigu', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Zuri', 'Zaizu', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Zuei', 'Zaizue', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Haiei', 'Zaie', 'Orainaldia');

Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Niri', 'Zaizkit', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Hari', 'Zaizkio', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Guri', 'Zaizkigu', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Zuri', 'Zaizkizu', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Zuei', 'Zaizkizue', 'Orainaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Haiei', 'Zaizkie', 'Orainaldia');

Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Niri', 'Zitzaidan', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Hari', 'Zitzaion', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Guri', 'Zitzaigun', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Zuri', 'Zitzaizun', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Zuei', 'Zitzaizuen', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Hura', 'Haiei', 'Zitzaien', 'Lehenaldia');

Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Niri', 'Zitzaizkidan', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Hari', 'Zitzaizkion', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Guri', 'Zitzaizkigun', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Zuri', 'Zitzaizkizun', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Zuei', 'Zitzaizkizuen', 'Lehenaldia');
Insert into nor_nori (nor, nori, aditz_lagunzailea, denbora) values ('Haiek', 'Haiei', 'Zitzaizkien', 'Lehenaldia');


Create table nor_nori_nork (
    nor TEXT,
    nori TEXT,
    nork TEXT,
    denbora TEXT,
    aditz_lagunzailea TEXT,
    PRIMARY KEY (nor, nori, nork, denbora),
    FOREIGN KEY (nor) REFERENCES nor(nor),
    FOREIGN KEY (nori) REFERENCES nori(nori),
    FOREIGN KEY (nork) REFERENCES nork(nork),
    FOREIGN KEY (denbora) REFERENCES denbora(denbora)
);

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Hari', 'Nik', 'Diot', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuri', 'Nik', 'Dizut', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuei', 'Nik', 'Dizuet', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Haiei', 'Nik', 'Diet', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Hari', 'Nik', 'Dizkiot', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuri', 'Nik', 'Dizkizut', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuei', 'Nik', 'Dizkizuet', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Haiei', 'Nik', 'Dizkiet', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Niri', 'Hark', 'Dit', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Hari', 'Hark', 'Dio', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Guri', 'Hark', 'Digu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuri', 'Hark', 'Dizu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuei', 'Hark', 'Dizue', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Haiei', 'Hark', 'Die', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Niri', 'Hark', 'Dizkit', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Hari', 'Hark', 'Dizkio', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Guri', 'Hark', 'Dizkigu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuri', 'Hark', 'Dizkizu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuei', 'Hark', 'Dizkizue', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Haiei', 'Hark', 'Dizkie', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Hari', 'Guk', 'Diogu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuri', 'Guk', 'Dizugu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuei', 'Guk', 'Dizuegu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Haiei', 'Guk', 'Diegu', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Hari', 'Guk', 'Dizkiogu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuri', 'Guk', 'Dizkizugu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuei', 'Guk', 'Dizkizuegu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Haiei', 'Guk', 'Dizkiegu', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Niri', 'Zuk', 'Didazu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Hari', 'Zuk', 'Diozu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Guri', 'Zuk', 'Diguzu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Haiei', 'Zuk', 'Diezu', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Niri', 'Zuk', 'Dizkidazu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Hari', 'Zuk', 'Dizkiozu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Guri', 'Zuk', 'Dizkiguzu', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Haiei', 'Zuk', 'Dizkiezu', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Niri', 'Zuek', 'Didazue', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Hari', 'Zuek', 'Diozue', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Guri', 'Zuek', 'Diguzue', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Haiei', 'Zuek', 'Diezue', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Niri', 'Zuek', 'Dizkidazue', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Hari', 'Zuek', 'Dizkiozue', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Guri', 'Zuek', 'Dizkiguzue', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Haiei', 'Zuek', 'Dizkiezue', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Niri', 'Haiek', 'Didate', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Hari', 'Haiek', 'Diote', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Guri', 'Haiek', 'Digute', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuri', 'Haiek', 'Dizute', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Zuei', 'Haiek', 'Dizuete', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Hura', 'Haiei', 'Haiek', 'Diete', 'Orainaldia');

Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Niri', 'Haiek', 'Dizkidate', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Hari', 'Haiek', 'Dizkiote', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Guri', 'Haiek', 'Dizkgute', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuri', 'Haiek', 'Dizkizute', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Zuei', 'Haiek', 'Dizkizuete', 'Orainaldia');
Insert into nor_nori_nork (nor, nori, nork, aditz_lagunzailea, denbora) values ('Haiek', 'Haiei', 'Haiek', 'Dizkiete', 'Orainaldia');

Create view aditz_lagunak(forme, nor, nori, nork, aditz_laguntzilea, denbora) as 
select 'nor_nork', nor, '', nork, aditz_lagunzailea, denbora from nor_nork
union
select 'nor_nori', nor, nori, '', aditz_lagunzailea, denbora from nor_nori
union
select 'nor_nori_nork', nor, nori, nork, aditz_lagunzailea, denbora from nor_nori_nork;


COMMIT;
