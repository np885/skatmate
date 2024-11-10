Was will ich bauen?
- Skatmate soll eine Übersicht aller gespielten Spiele in einem Webbrowser darstellen
	- Skatrunden importieren und persistieren (csv, Format?)
	- Skatrunden ausgeben in Klassischer Notation und Bierlachs (inkl. gewinner)
	
Womit will ich es bauen?
- Serverside
	- AWS Lambda -> Java
- Client
	- AWS S3 -> Typescript/JS
- Database 
	- MongoDB

DevOps Themen:
	- Terraform zum erstellen der AWS ressourcen
	- CI/CD Pipeline mit Github

Was sind die nächsten Schritte?
1. Technischen Durchstich Infrastruktur
   - AWS Ressourcen mit Terraform hinstellen
      - S3 Bucket 
      - API Gateway 
      - lambda function
   - Hello World Website
     - simple java lambda mit String return
     - simple website mit ajax call to lambda
2. Lambda fürs lesen der Skatrunden erstellen
   - Lesen einer skatrunde aus dem backend
   - Im Frontend anzeigen
3. Database anbinden
   - MongoDB Zugriff in stack integrieren
   - Persistenz in lambda anbinden
   - Beispieldaten in Mongo DB hinterlegen
4. Frontend refactor
   - React einbinden
   - UI Lib einbinden
5. Import lambda für Schnittstellen bauen