# **Robot Management Service**

Robot Management Service adalah REST API yang dibuat menggunakan bahasa pemrograman **Go** dan framework **Gin** untuk mengelola data robot. 

## **Cara Menjalankan Proyek**

### **1. Persyaratan Sistem**
- **Go**: Pastikan Go telah terinstal 
- **Postman** atau alat HTTP lainnya untuk mengakses API.

### **2. Clone Repository**
```bash
git clone https://github.com/elsafir/test-backend-maxchat.git
cd <repository_folder>
```
### **3. Jalankan Proyek**
1. Pastikan Anda berada di direktori yang berisi file `main.go` dan `data.txt`.
2. Jalankan perintah berikut:
   ```bash
   go run main.go
   ```
3. Server akan berjalan di `http://localhost:8080`.

---

## **Endpoint API**(Gunakan Postman)

### **1. Mendapatkan Semua Data Robot**
- **URL**: `/robots`
- **Metode**: `GET`
- **Response**:
  ```json
  [
      {
          "code": "rv1",
          "name": "Rover #1",
          "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
          "model": "car",
          "tech": ["AI", "car"],
          "status": "progress"
      },
      ...
  ]
  ```

### **2. Mendapatkan Data Robot Berdasarkan Kode**
- **URL**: `/robots/:code`
- **Metode**: `GET`
- **Parameter Path**: `:code` (contoh: `rv1`)
- **Response** (contoh):
  ```json
  {
      "code": "rv1",
      "name": "Rover #1",
      "description": "Lorem ipsum dolor sit amet",
      "model": "car",
      "tech": ["AI", "car"],
      "status": "progress"
  }
  ```

### **3. Menambahkan Robot Baru**
- **URL**: `/robots`
- **Metode**: `POST`
- **Body** (JSON):
  ```json
  {
      "code": "new1",
      "name": "New Robot",
      "description": "A new robot description",
      "model": "car",
      "tech": ["AI", "robot"],
      "status": "active"
  }
  ```
- **Response**:
  ```json
  {
      "code": "new1",
      "name": "New Robot",
      "description": "A new robot description",
      "model": "car",
      "tech": ["AI", "robot"],
      "status": "active"
  }
  ```

### **4. Memperbarui Robot**
- **URL**: `/robots/:code`
- **Metode**: `PUT`
- **Parameter Path**: `:code` (contoh: `rv1`)
- **Body** (JSON):
  ```json
  {
      "name": "Updated Robot",
      "description": "Updated description",
      "model": "car",
      "tech": ["AI"],
      "status": "active"
  }
  ```
- **Response**:
  ```json
  {
      "code": "rv1",
      "name": "Updated Robot",
      "description": "Updated description",
      "model": "car",
      "tech": ["AI"],
      "status": "active"
  }
  ```

### **5. Menghapus Robot**
- **URL**: `/robots/:code`
- **Metode**: `DELETE`
- **Parameter Path**: `:code` (contoh: `rv1`)
- **Response**:
  ```json
  {
      "message": "Robot deleted"
  }
  ```

### **6. Filter Robot**
- **URL**: `robots/filters`
- **Metode**: `GET`
- **Query Parameters**:
  - `model` (opsional): Filter berdasarkan model (contoh: `humanoid`).
  - `tech` (opsional, multi-filter): Gunakan beberapa parameter untuk filter multi-teknologi (contoh: `tech=AI&tech=robot`).
- **Response**:
  ```json
  [
      {
          "code": "px1",
          "name": "Pacifista 1",
          "description": "Lorem ipsum dolor sit amet",
          "model": "humanoid",
          "tech": ["AI", "robot"],
          "status": "active"
      }
  ]
  ```

---
