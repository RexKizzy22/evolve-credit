BEGIN;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    date DATE NOT NULL,
    country VARCHAR NOT NULL
);

INSERT INTO users (name, email, date, country) 
VALUES ('Yvonne Haynes','lobortis.mauris.suspendisse@protonmail.ca','Dec 15, 2021','Costa Rica'),
        ('Jerry Mccarty','suspendisse@aol.edu', 'Dec 9, 2021', 'Turkey'),
        ('Stephen Coffey','molestie.sed@icloud.ca','Dec 12, 2021','Spain'),
        ('Madison Orr','morbi@google.ca','Jan 17, 2022','Italy'),
        ('Maite Klein','imperdiet.nec@yahoo.couk','May 20, 2023','Australia'),
        ('Sonya Wheeler','dolor.nulla@icloud.ca','Mar 20, 2022','Turkey'),
        ('Nola Duke','augue.malesuada@google.edu','Nov 4, 2022','Ukraine'),
        ('Scarlett Weeks','aenean@aol.net','Feb 5, 2023','New Zealand'),
        ('Walter Roberts','aliquam.erat.volutpat@aol.ca','Dec 19, 2021','Pakistan'),
        ('Rudyard Dudley','dignissim.lacus@yahoo.net','Apr 15, 2023','Pakistan'),
        ('Nasim Austin','malesuada@google.couk','Jan 7, 2022','Vietnam'),
        ('Jayme Hansen','quam@outlook.org','Apr 20, 2023','Belgium'),
        ('Mara Mcclure','non@protonmail.net','Nov 23, 2021','Colombia'),
        ('Celeste Mckinney','mauris.eu.turpis@aol.org','Mar 25, 2023','Canada'),
        ('Alma Joyce','massa@aol.ca','Jun 17, 2022','United Kingdom'),
        ('Susan Cook','vestibulum@hotmail.ca','Feb 28, 2022','Colombia'),
        ('Mannix Huff','est.congue@protonmail.net','Dec 29, 2022','Australia'),
        ('Indigo Marks','aliquam.tincidunt@google.ca','Dec 5, 2022','China'),
        ('Burton Odom','in.faucibus@aol.com','Dec 25, 2021','Colombia'),
        ('Tara Owen','mattis@google.edu','Nov 15, 2022','New Zealand'),
        ('Hoyt Alvarez','sit.amet.ante@aol.ca','Apr 23, 2023','Vietnam'),
        ('Olivia Brown','congue.turpis@google.couk','Nov 25, 2021','Canada'),
        ('Jeanette Preston','aliquam.ornare@icloud.edu','Jun 2, 2022','Philippines'),
        ('Lisandra Boyer','diam.nunc.ullamcorper@outlook.net','Jan 8, 2023','Germany'),
        ('Brett Vang','iaculis.enim.sit@hotmail.edu','Jul 21, 2021','Belgium'),
        ('Katell Atkinson','etiam.ligula@protonmail.org','Jul 24, 2022','Peru'),
        ('Tasha Vance','neque@icloud.couk','Jul 3, 2022','Vietnam'),
        ('Isadora Ross','amet@protonmail.com','Jan 27, 2023','Netherlands'),
        ('Branden Oneal','quisque.nonummy@yahoo.ca','Feb 21, 2023','Peru'),
        ('Hilel Hernandez','nisl.elementum@protonmail.net','Apr 19, 2022','Peru');

COMMIT;