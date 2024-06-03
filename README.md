# Hospital Management System

This is a README file for the Hospital Management System built using Go and Gin, PostgreSQL, and hosted on Google Cloud
Platform.

## Getting Started

To start the server, run the following command:

```sh
go run main.go
```

## Endpoints

### Patient

- **GET** `/patients`
    - Retrieves a list of patients.
      Example: [https://hospital-management-424616.ue.r.appspot.com/patients](https://hospital-management-424616.ue.r.appspot.com/patients)

- **GET** `/patients/search?name=jane&phone=98&email=jane`
    - Searches for a patient by name, phone, and email.
      Example: [https://hospital-management-424616.ue.r.appspot.com/patients/search?name=jane&phone=98&email=jane](https://hospital-management-424616.ue.r.appspot.com/patients/search?name=jane&phone=98&email=jane)

### Medical Record

- **GET** `/medical-record/doctor/434a4cbe-53a3-4945-9a0a-d1a05d02cef8/report?start_date=2024-06-02&end_date=2024-06-05`
    - Retrieves medical records for a doctor within a specified date range.
  
      Example: [https://hospital-management-424616.ue.r.appspot.com/medical-record/doctor/434a4cbe-53a3-4945-9a0a-d1a05d02cef8/report?start_date=2024-06-02&end_date=2024-06-05](https://hospital-management-424616.ue.r.appspot.com/medical-record/doctor/434a4cbe-53a3-4945-9a0a-d1a05d02cef8/report?start_date=2024-06-02&end_date=2024-06-05)

- **PATCH** `/medical-record/patient/fb4b6625-439a-4131-a206-9e61389b0cb0`
    - Updates a patient's medical record with a diagnosis.
    - Content-Type: application/json
  ```json
  {
    "diagnosis": "Acute encephalitis"
  }
  ```
  Example: [https://hospital-management-424616.ue.r.appspot.com/medical-record/patient/fb4b6625-439a-4131-a206-9e61389b0cb0](https://hospital-management-424616.ue.r.appspot.com/medical-record/patient/fb4b6625-439a-4131-a206-9e61389b0cb0)

### Medical Appointment

- **GET** `/medical-appointments/doctor/c8e10804-bbce-49f6-a73f-ce80629c2b23`
    - Retrieves medical appointments for a specific doctor.
  
      Example: [https://hospital-management-424616.ue.r.appspot.com/medical-appointments/doctor/c8e10804-bbce-49f6-a73f-ce80629c2b23](https://hospital-management-424616.ue.r.appspot.com/medical-appointments/doctor/c8e10804-bbce-49f6-a73f-ce80629c2b23)

- **POST** `/medical-appointments`
    - Creates a new medical appointment.
    - Content-Type: application/json
  ```json
  {
    "patient_id": "aa1758e0-dde6-4071-b18a-022036762968",
    "doctor_id": "c8e10804-bbce-49f6-a73f-ce80629c2b23",
    "date_time": "2024-10-03 13:00:00",
    "reason": "Urgent appointment"
  }
  ```
  Example: [https://hospital-management-424616.ue.r.appspot.com/medical-appointments](https://hospital-management-424616.ue.r.appspot.com/medical-appointments)

## About the Application

The Hospital Management System is a web application designed to streamline various aspects of hospital management,
including patient information management, medical records management, and appointment scheduling. Built using Go and Gin
framework, PostgreSQL for database management, and hosted on Google Cloud Platform, this system ensures reliability,
scalability, and security for hospital operations.