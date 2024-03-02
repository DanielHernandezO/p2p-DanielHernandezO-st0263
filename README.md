# P2P Para Descaga De Archivos - configuracion
**Nota**: La descripcion, arquitectura del proyecto y cada uno de los servicios se puede encontrar en la [Wiki](https://github.com/DanielHernandezO/p2p-DanielHernandezO-st0263/wiki/P2P--para-compartir-archivos) del proyecto, donde igualmente se incluye los aspectos solicitados para la entrega.

Para la configuracion de un peer se debe configurar las properties de cada uno de los servicios de acuerdo a las siguientes variables.

###  client 

Para configurar este servicio encontramos un archivo `aplication.properties` ubicado `client/configs/` que cuenta con las siguientes variables:

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `ip_download`     | `string` | **Required**. Es la IP donde esta corriendo el servicio del download, para este caso la misma del peer |
| `port_download`     | `string` | **Required**. Es el pueto donde esta corriendo el servicio del download **cuidado: debe ser el mismo configurado en el servicio del download**|
| `ip_server`     | `string` | **Required**. IP donde estara corriendo el search que para este caso seria la misma ip del peer|
| `port_server`     | `string` | **Required**. Puerto donde esta corriendo el servicio del search **cuidado: debe ser el mismo configurado en el servicio del search**|
| `peers`     | `map<string,string>` | **Required**. IP de los peers a los que esta conectado y el puerto donde esta corriendo el servicio del search en el peer|
| `base_filemanagment_url`     | `string` | **Required**. Url base del servicio del file_management|
| `filemanagment_post_file`     | `string` | **Required**. Path usado para hacer post de los archivos en el servicio del file_management|
| `ip_filemanagment`     | `string` | **Required**. IP donde estara corriendo el servicio del file_management que para este caso seria la misma ip del peer|
| `port_filemanagment`     | `string` | **Required**. Puerto donde esta corriendo el servicio del file_management **cuidado: debe ser el mismo configurado en el servicio del file_management**|
| `files_location`     | `string` | **Required**. Directorio donde estan ubicados los archivos que van a estar disponibles en el peer|

###  search 

Para configurar este servicio encontramos un archivo `aplication.properties` ubicado `search/configs/` que cuenta con las siguientes variables:

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `network`     | `string` | **Required**. Protocolo de transmision usado **para este caso tcp**|
| `ip_host`     | `string` | **Required**. IP donde estara corriendo el servicio del search que para este caso seria la misma del peer|
| `port_host`     | `string` | **Required**. puerto configurado donde va a correr el servicio del search|
| `base_filemanagment_url`     | `string` | **Required**. url base del servicio del file_management|
| `filemanagment_get_file`     | `string` | **Required**. Path usado para hacer get de los archivos en el servicio del file_managemen|
| `ip_filemanagment`     | `string` | **Required**. IP donde estara corriendo el servicio del file_management que para este caso seria la misma ip del peer|
| `port_filemanagment`     | `string` | **Required**. Puerto donde esta corriendo el servicio del file_management **cuidado: debe ser el mismo configurado en el servicio del file_management**|
| `base_rabbit_url`     | `string` | **Required**. Url base para conectarse a rabbitMQ|
| `user_rabbit`     | `string` | **Required**. Usuario configurado para conectarse a rabbitMQ|
| `password_rabbit`     | `string` | **Required**. Password configurado para conectarse a rabbitMQ|
| `ip_rabbit`     | `string` | **Required**. IP donde estara corriendo el servicio del rabbitMQ que para este caso seria la misma del peer|
| `port_rabbit`     | `string` | **Required**. Puerto configurado donde va a correr rabbitMQ|
| `queue_download`     | `string` | **Required**. Nombre de la cola donde se van a publicar los mensaje **cuidado: dede ser la misma donde esta consumiendo el download**|
| `exchange_download`     | `string` | **Required**. Nombre del exchange usado para la cola**|

###  Download 

Para configurar este servicio encontramos un archivo `aplication.properties` ubicado `Download/configs/` que cuenta con las siguientes variables:

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `port`     | `string` | **Required**. Puerto donde va a correr el servicio |
| `base_download_url`     | `string` | **Required**. Url base para la transferencia de archivos entre servicios Download|
| `base_rabbit_url`     | `string` | **Required**. Url base para conectarse a rabbitMQ|
| `user_rabbit`     | `string` | **Required**. Usuario configurado para conectarse a rabbitMQ|
| `password_rabbit`     | `string` | **Required**. Password configurado para conectarse a rabbitMQ|
| `ip_rabbit`     | `string` | **Required**. IP donde estara corriendo el servicio del file_management que para este caso seria la misma del peer|
| `port_rabbit`     | `string` | **Required**. Puerto configurado donde va a correr rabbitMQ|
| `queue_download`     | `string` | **Required**. Nombre de la donde se van a consumir los mensaje **cuidado: dede ser la misma donde se esta publicando**|
| `files_location`     | `string` | **Required**. Ubicacion donde se van a guardar los archivos descargados **|

###  Fetch 

Para configurar este servicio encontramos un archivo `aplication.properties` ubicado `fetch/configs/` que cuenta con las siguientes variables:

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `ip`     | `string` | **Required**. IP donde estara corriendo el servicio del fetch que para este caso como estan en el mismo peer seria la misma del peer|
| `port`     | `string` | **Required**. Puerto donde esta corriendo el servicio del search **cuidado: debe ser el mismo configurado en el servicio del search**|
| `time`     | `string` | **Required**. Tiempo en milisegundos de cada cuanto tiempo se va a ejecutar el fetch|

###  file_management 

Para configurar este servicio encontramos un archivo `aplication.properties` ubicado `file_management/configs/` que cuenta con las siguientes variables:

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `port`     | `string` | **Required**. Puerto configurado donde va a correr el servicio del file_management|
| `files_location`     | `string` | **Required**. Ubicacion donde se van a almacenar los archivos que el cliente publique|


##  Requerimientos del sistema
- python3(pip, grpcio, python-dotenv)
- go 1.22
- docker
- rabbitMQ

##  ¿Como correr el peer?
Para correr cada peer necesitamos:
- Crear una instancia de EC2 con OS Ubuntu 22.04
- Clonar el proyecto con el comando `git clone https://github.com/DanielHernandezO/p2p-DanielHernandezO-st0263.git`
- ubicarse en la carpte `p2p-DanielHernandezO-st0263` con el comando `cd p2p-DanielHernandezO-st0263`
- Asignar permisos de ejecucion al archivo con el comando `chmod +x run.sh`
- Correr el archivo con el siguiente comando `./run.shell` que inicialmente validara que cada uno de los requisitos mencionados anteriormente se cumplan, en caso que no se cumplan, instalara las herramientas mencionadas y creara una imagen de docker para rabbitMQ con el nombre `rabbit-server` con el user y el password igual a `default` si alguno de estos valores se desea modificar se puede hacer directamente en el script. Lugeo se correra cada uno de los servicios que conforman el servidor del peer y se va almacenar un archivo `.log` en la carpeta `logs` por cada servicio.
- En caso de no tener instalado en go en el sistema anteriormente se debe correr este comando en consola `source ~/.bashrc`
- Por ultimo, se debe ingresar a la carpeta `p2p-DanielHernandezO-st0263/client` y ejecutar el cliente con el comando `go run cmd/app/main.go`


## Autor

- [@DanielHernandezO](https://www.github.com/DanielHernandezO)


