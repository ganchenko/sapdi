Data Extraction from SAP S/4HANA CDS View to a File Store
===========
#### Description
The Graph extracts data fom an ABAP CDS View to a file store and creates related files in the file store; 
specifically in this example, the source is an SAP S/4HANA.
Not mandatory, but for convenience, this example may be extended to read the data from the written files and to feed into the Terminal operator.

#### Prerequisites
1. You need a valid connection to SAP S/4HANA. 
   Please create a valid and checked connection to an SAP S/4HANA environment or to an SAP NetWeaver/ABAP based system which supports the CDS technology.
   Please use the related description of the SAP DI - Connection Management to implement and test such a connection.
   In this example the communication is based on HTTPS, but you may use other protocols such as HTTP, RFC or (WebSocket) WSRFC
2. You need a valid connection to a file store and the ability to create, read and write files to that file store location.
   Please use the related description of the SAP DI - Connection Management to implement and test such a connection.
3. A couple of Operators need to be in place: (ABAP) CDS Reader, ABAP Converter, Go Operator (--> Determine Last Batch, Limit File Size, Check Last Batch), Write(Read) File Operator, Graph Terminator
4. The CDS View must allow data extraction via the annotation: "@Analytics:{ dataExtraction: enabled: true ..." - for delta extraction also the "delta.changeDataCapture" part of the annotation needs to be maintained correctly.


#### Configure and Run the Graph

1. In this example the ABAP CDS Reader is configured to execute an initial load.
   The CDS Reader supports 3 ways to load data:
   a) the straight forward Initial Load, which copies the data from source to target but without any subsequent replication.
   b) the Replication, which includes the Initial Load and - based on a delta handling - the subsequent replication of changed data.
   c) the Delta Processing, which implies that no Initial Load takes place and changed data only will get processed.
2. The ABAP Converter feeds via a Message Converter into a Go Operator 'Determine Last Batch', which checks whether the SLT-based replication got stopped and is then launching a soft-kill (used at 'Check Last Batch')
3. The data keep flowing into a Go Operator 'Limit File Size', which controls the file size as a parameter and the file counter.
4. The Write File Operator is used to create the files until end-of-initial-load.
   The Write File Operator uses the configured 'append' mode to fill into `/tmp/file_<header:ABAPfilenumber>.csv` and its configured connection.
5. The Go Operator 'Check Last Batch' terminates the graph in case of the end of the replication; 
   without the two operators '* Last Batch', it can happen that the graph gets terminated before the last write got completed.
6. (optional) Simply to visualize the file, the just used and created file may be read again and sent to Terminal.
7. Use <RUN> to execute this Graph.
   Note: in case of an Initial Load only, the graph will terminate, in case of Delta Mode or Replication (=Initial Load plus Delta) the graph will not stop after the initial load.

#### Data Integration and Data Security
Please make yourself familiar with the related Integration Guides and the Security considerations as of
the SAP Data Hub "ABAP Integration Guide" https://help.sap.com/doc/61c7b7e293b74a45b50724c285df9560/2.7.latest/en-US/loio61c7b7e293b74a45b50724c285df9560.pdf ,
the SAP Data Hub "Administration Guide" https://help.sap.com/viewer/3f4043064eed446a895bc8ba7e61dc83/2.7.latest/en-US , and
SAP Note 2831756 "SAP Data Hub/ Data Intelligence ABAP Integration - Security Settings" 

<br>
<div class="footer">
   &copy; 2020-2022 SAP SE or an SAP affiliate company. All rights reserved.
</div>
