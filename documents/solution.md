Components Overview
-------------------
- **Cloud IoT Core**: Securely connects and manages power meters at scale.
- **Pub/Sub**: Provides real-time messaging and data ingestion, acting as a buffer and managing data flow.
- **Dataflow**: Used to processing, transforming, and enriching data before storage.
- **BigQuery**: Data warehousing, storage, analysis, and querying of large datasets.
- **Cloud Storage**: Cost effective component for raw data backup
- **Identity and Access Management (IAM)**: Secures access to data and resources.
- **Cloud Run**: Used hosting customer-facing applications that interact with the backend.
- **Apigee**: Add security, access control, rate limiting and logging to Cloud Run.
- **Terraform**: Automates the deployment and management of infrastructure resources.

Data Flow and Architecture
---------------------------
**Data Collection**: Each power meter transmits usage data to Cloud IoT Core.

**Data Ingestion**: Data from IoT Core is published to a Pub/Sub topic, providing a scalable entry point for real-time data.

**Processing**: Dataflow subscribes to the Pub/Sub topic, processing the data in real-time or batches. This stage involves validating, transforming, and enriching the data (building and customer information).

**Data Storage**:
- *Storage*: Processed data is then loaded into BigQuery, with tables structured to support use cases such as monthly energy consumption, building comparison, and data analysis.
- *Glacial Storage*: Raw data / backups can be stored in Cloud Storage for long-term storage or future processing.

**Data Analysis and Access**:
- Customers access their data through applications or APIs hosted on Cloud Run, which interact with BigQuery to fetch and display energy usage data.
- BigQuery allows for complex queries and analysis.

Considerations
--------------
- **Scalability**: GCP services like Pub/Sub and BigQuery are designed to scale automatically.
- **Reliability**: Managed services ensure high availability and reliability of the data ingestion and processing pipeline.
- **Security**: Cloud IoT Core and IAM provide security mechanisms to authenticate devices and control access to data.
- **Cost**: Pay-as-you-go pricing models of GCP allow for cost-effective scaling
