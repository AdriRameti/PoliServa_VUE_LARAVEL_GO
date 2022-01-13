# POLISERVA
Realizado por: 
Adrián Ramos Ureña y Cristian Juan Saez

## Tecnologías y sus ámbitos de uso
La aplicación que veremos a continuación ha sido desarrollada con Vue 3 , Laravel ,Go y Bootstrap 5.1. Seguidamente, explicamos para cada ambito de uso, que tecnología hemos empleado.

### Tecnologías
[Bootstrap 5.1](https://getbootstrap.com/docs/5.1/getting-started/introduction/ "Bootstrap")
[VUE 3](https://v3.vuejs.org/ "Vue")
[Laravel](https://laravel.com/ "Laravel")
[Eloquent ORM](https://laravel.com/docs/8.x/eloquent "Eloquent")
[GO](https://go.dev/ "GO")
[Gorm ORM](https://gorm.io/index.html "Gorm")

### Diseño
Para el diseño de la aplicación, hemos utilizado Bootstrap 5.1 para aplicar todos los conocimientos vistos en clase. Posteriormente, en la demostración de la aplicación, explicaremos que elementos hemos utilizado en cada página.

### Cliente o Frontend 
En el desarrollo de la parte cliente de la aplicación, hemos utilizado Vue 3 junto a Vuex y Vue Router, que parten del framework progresivo de Vue. A diferencia de otros clientes con framework monoliticos, Vue hace referencia a un framework híbrido ya que parte de una libreria propia central, pero se le pueden añadir distintos paquetes adicionales. Este cliente será el encargado de realizar todas las operaciones y peticiones a nuestros servidores. Como hemos dicho anteriormente, posteriormente explicaremos que elementos hemos utilizado en cada página.

### Servidor o Backend
En el desarrollo de la parte servidora, hemos utilizados dos servidores distintos los cuales son Laravel (con su propio ORM Eloquent) y GO (con su propio ORM Gorm). 
El servidor Laravel ha sido destinado a gestionar las migraciones de la base de datos y todas las operaciones relacionadas con los usuarios.
El servidor Go ha sido destinado a gestionar todas las operaciones simples de obtención de datos , crear nuevos registros o modificarlos.

## Demostración de la aplicación

### Home
 El Homepage será la vista principal de la aplicación y será donde nos encontremos un carousel (hecho con Bootstrap) y el apartado de categorias donde estarán todos los deportes (hecho con la clase cards de Bootstrap). 
 En lo que hace referencia a Vue, en el apartado Home se ha utilizado una unica consulta en el setup donde obtendremos todos los deportes y recorremos el resultado con un v-for para que vaya añadiendo categorias con relación a los deportes. La consulta esta hecha desde un composable de vue y accede al servidor de Go donde obtenemos todos los deportes, para hacer uso de los composables en el setup, hemos hecho un async setup, el qual, para que funcione hemos hecho utilidad del componente Suspense de Vue.
 
 ![image](https://user-images.githubusercontent.com/75810680/149213451-978bac7a-58cf-4ec8-8e1c-b7ed0945a721.png)
 
 ![image](https://user-images.githubusercontent.com/75810680/149213509-7c9047f7-3fef-41a7-a702-13500e174b62.png)
 
 ### Reservas
 La vista de reservas será la que nos muestre las pistas de los deportes. A esta vista podremos acceder desde la barra de navegación donde nos mostrará todas las pistas (utilizando la clase Cards de Bootstrap) de todos los deportes, o podremos acceder a ella haciendo click encima de una categoria de deporte en la vista Home, donde nos mostrará las pistas libres de ese deporte. 
 
 Dentro de esta vista, podremos filtrar las pistas por la fecha (utilizando un datepicker de vue) y hora (si elegimos por ejemplo las 15:00 en la hora inicial, en la hora final se bloquearan todas las horas que sean iguales o inferiores). Hasta que no apliquemos el filtro de fecha y hora no se desbloqueará el botón de confirmar reserva, el cual al pulsarlo, nos aparecerá un modal (hecho con Bootstrap) donde nos mostrará la información de la reserva. Al confirmar la reserva desde el modal, se nos guardará la reserva y nos redirigirá al Homepage, y en caso de que no estemos logueados, nos redigirá al Login. Dependiendo de las operaciones que hagamos, tendremos mensajes de exito o error ejecutados a partir de un Toastr.

En lo que hace referencia a Vue, hay 3 posibles consultas en el setup dependiendo del redireccionamiento que se haya elegido. Cada consulta esta realizada en un composable y irá redireccionada al servidor de Go (todas las pistas, las pistas del deporte seleccionado o luego de filtrar por fecha y hora de reserva). Además se utilizan las directivas v-for para recorrer el resultado e ir añadiendo las pistas, v-show o v-if para mostrar unos elementos o otros ... 
 Por otra parte, la petición de la reserva se realizará a Laravel donde obtendrá el usuario actual y se comunicará con Go una vez ya tenga toda la información (la de la reserva más la del usuario) para realizar la reserva.
 
 ![image](https://user-images.githubusercontent.com/75810680/149213888-8dc20ca4-7f1e-4e26-89a5-f0000b242d10.png)
 
 ![image](https://user-images.githubusercontent.com/75810680/149213941-c2ed38cf-d353-43d9-9ff0-c9c9851f4805.png)
 
 ![image](https://user-images.githubusercontent.com/75810680/149213812-d306eecc-eb7f-43e8-85e9-efd6bb0e650d.png)
 
 ![image](https://user-images.githubusercontent.com/75810680/149214020-1e89aaad-7609-40aa-8744-e64eac778eca.png)
 
 ### Register
 En la vista del register, nos mostrará un formulario (realizado con Bootstrap), donde nos pedirán la información necesaria para realizar el registro del usuario. 
 
 En lo que se refiere a Vue, obtendremos los valores del formulario a partir de variables declaradas en data y obtenidas con el v-model. Además utilizaremos el watch para realizar una barra de seguridad de la contraseña (realizado con Bootstrap) dependiendo de la longitud de la contraseña. También validaremos todos los campos del formulario para que sean correctos a la hora de realizar el registro.
 Las peticiones se realizarán al store de Vue que se comunicará con Laravel. Habrá dos peticiones al servidor ya que antes de registrar al usuario, se le enviará un correo con un código generado aleatoriamente el cual deberá poner para verificar que es la persona que se va registrar. Una vez verificado el código, se realizará el registro del usuario.
 
![image](https://user-images.githubusercontent.com/75810680/149214163-efa1ab69-d930-45a6-850b-688c87481626.png)
 
 ### Login 
 En la vista del login, nos mostrará un formulario (realizado con Bootstrap), donde a partir del correo y contraseña utilizado en el register, podra acceder con su cuenta.
 En lo que se refiere a Vue, obtendremos los valores del formulario a partir de variables declaradas en data y obtenidas con el v-model. Realizaremos una verificación donde si hay algun error nos avisará un toastr con el problema, que el usuario no exista o que la contraseña es incorrecta, por ejemplo.
 
 Esta vista solo tendrá una consulta que se realizará al store de Vue y se comunicará con laravel para obtener la información del usuario. En caso de que el usuario tenga la verificación de dos pasos, también implementada en el servidor de Laravel, una vez puesto el correo y contraseña, en caso de ser correcto, nos redigirá a una vista donde tendremos que poner el código de la aplicación de verificación.
 
![image](https://user-images.githubusercontent.com/75810680/149214247-76f5b798-82c7-4be5-adc7-6bd18c99de4b.png)
 
 ### Profile 
 En la vista del profile, nos aparecerá las estadisticas(utilizando la libreria de vue-chartkick) de cuantas reservas se han realizado en el último mes y que deportes han sido los que ha reservado el usuario. Además, también habrá un apartado donde podremos modificar la información del usuario por si queremos modificar el nombre, apellidos, la contraseña, el correo o la foto del perfil. También se podrá activar la verificación en dos pasos, donde escanearemos el codigo QR y se nos guardará el correo junto con el código en la propia aplicación. Todo el diseño, modales y distintas vistas del profile se han realizado con Bootstrap.
 
 En lo que se refiere a Vue, tendremos en el setup las dos peticiones realizadas a Go utilizando composables para obtener y mostrar la información de las gráficas. Además utilizaremos distintos métodos en la activación de la verificación de dos pasos, para realizar la modificación de los datos del usuario (Obteniendo los datos con v-model), para activar y desactivar variables que iran asociadas a v-show y v-if etc.
 
 ![image](https://user-images.githubusercontent.com/75810680/149214551-3ea7f157-681d-4206-84cc-43b78552cbef.png)
 
 ![image](https://user-images.githubusercontent.com/75810680/149214598-53a381bc-d27a-4813-9289-2fd60094ed43.png)
 
 ### Dashboard
 La vista del Dashboard solo se mostrará si el usuario tiene los permisos de administrador (utilizando guards en Vue). Si es así, le aparecerá una nueva opción en la barra de navegación desde la cual podra acceder al dashboard del administrador. En él, podrá observar la estadistícas globales de las reservas, por deportes y por pistas; ver la información de todas las pistas, de todos los deportes e incluso de todos los usuarios. El administrador podrá añadir nuevos elementos a los distintos apartados, modificarlos o incluso eliminarlos. Pero siempre verificando que es el administrador y tiene los permisos.
 
 En lo que se refiere a Vue, las estadísticas utilizarán la libreria de vue-chartkick, y la información de cada apartado de administración se mostrará con datatables utilizando las librerias (ag-grid-community y ag-grid-vue3).
 
![image](https://user-images.githubusercontent.com/75810680/149214683-a7b7e2b9-53e3-4197-a6dd-8257f99dbf48.png)

![image](https://user-images.githubusercontent.com/75810680/149214754-386777ba-a00c-490a-a8b0-40e8cf53a154.png)
 
 ### Verificaciones
 Para las distintas verificaciones dentro de la aplicación, hemos implementado dos métodos distintos que se utilizarán dependiendo de la acción que se desee realizar.
 
 El primero es verificar a partir de un código que se envia al correo y el cual es generado aleatoriamente. Cuando se requiere de esta validación, hasta que no se introduce el codigo y se comprueba que el código introducido es el mismo que el que se envió, la operación no seguira adelante. Para esta verificación, hemos utilizado Mailgun, una api externa de envio de correos.
 
![image](https://user-images.githubusercontent.com/75810680/149214904-5a44e588-30e9-43fe-a1f1-484ee9752a93.png)
 
 El segundo metodo de verificación, es la verificación en dos pasos. Al activar esta verificación, se mostrará un QR que al escanearlo desde la aplicación de verificación, guardará el codigo, así cuando se precise de validaciones por 2 pasos, simplemente deberemos introducir el código que nos muestra la aplicación de verficación, luego se verificará y si es correcto seguirá con la operación.

![image](https://user-images.githubusercontent.com/75810680/149214993-499fa1c2-7670-4448-acff-f1e5921f48eb.png)

 
