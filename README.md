# InstaFix
InstaFix is a client-server desktop application made as a project for the course "Advanced Programming Language" @ University of Catania. It was made in collaboration with my
university colleague @PhilipTamb <br />
The aim of InstaFix is to put the user in contact with professionals in the area to solve everyday household problems: repairs, plumbing, cleaning, electrical system
problems, gardening, carpentry and many other services that require the intervention of a specialist. <br />
The architecture of this application is composed by 2 clients (one for the user and the other one for the professionists) and a server. The user client is written in C#,
the professionist client in Python and the server was made with Go Language. <br />
The server acts as an intermediary between the two clients and is the only entity that has direct access to the database; we have chosen to use a relation DB of the mySQL type <br />
<br />
PS the documentation is the pdf file in the "Documentation" folder, but it's in italian. There's also a demonstration video of the C# client
