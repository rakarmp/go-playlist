## Go-Playlist

Simple Playlist API, Pure GO

### **Available Param**

```
| Route                       | HTTP Method | Description             | URL Parameters | Body Parameters |
|-----------------------------|-------------|-------------------------|----------------|-----------------|
| `/playlist`                 | GET         | Find All Playlists      | -              | -               |
| `/playlist`                 | POST        | Create a New Playlist   | -              | Playlist Data   |
| `/playlist/:id`             | GET         | Find Playlist by ID     | `id` (Path)    | -               |
| `/playlist/:id`             | PATCH       | Update Playlist         | `id` (Path)    | Playlist Data   |
| `/playlist/:id`             | DELETE      | Delete Playlist         | `id` (Path)    | -               |
```
