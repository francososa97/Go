# API de Productos en Go

Esta es una API simple para la gestión de productos desarrollada en Go. Permite realizar operaciones CRUD (Crear, Leer, Actualizar, Eliminar) sobre productos.

## Características

- Crear un nuevo producto con nombre, descripción, precio y cantidad en stock.
- Obtener la lista de todos los productos.
- Obtener un producto específico por su ID.
- Actualizar la información de un producto existente.
- Eliminar un producto por su ID.

## Uso

### Instalación

1. Clona este repositorio:

```bash
git clone https://github.com/tu_usuario/api-productos-go.git
```

2. Clona este repositorio:
```
cd api-productos-go
```
3. Instala las dependencias:

```bash
go mod tidy
```
## Configuración
Configura la base de datos en el archivo config/config.go. Actualiza los valores de conexión según tu entorno:
```
const (
    DBHost     = "localhost"
    DBPort     = "5432"
    DBUser     = "tu_usuario"
    DBPassword = "tu_contraseña"
    DBName     = "nombre_base_de_datos"
)
```

## Ejecución
Compila y ejecuta la aplicación:
```
./api-productos-go
La API estará disponible en http://localhost:8080.
```


## Endpoints
- GET /productos: Obtiene la lista de todos los productos.
- GET /productos/{id}: Obtiene un producto específico por su ID.
- POST /productos: Crea un nuevo producto.
- PUT /productos/{id}: Actualiza la información de un producto existente.
- DELETE /productos/{id}: Elimina un producto por su ID.
