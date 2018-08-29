# 1.  Connecting to the service instance.

# Enable the required Python libraries.
import csv
from cloudant.client import Cloudant
from cloudant.error import CloudantException
from cloudant.result import Result, ResultByKey


# Useful variables
serviceUsername = "004ab76a-6ba7-4501-89d7-92e9b59a907a-bluemix"
servicePassword = "cefebb2749f81c39eb1062f3b412f3276cfbe9ad199d8f7894ce502fb8dacde8"
serviceURL = "https://004ab76a-6ba7-4501-89d7-92e9b59a907a-bluemix.cloudant.com"

# This is the name of the database we are working with.
databaseName = "my_code_call"

# This is a simple collection of data,
# to store within the database.
with open('fileRE.csv', 'rb') as f:
    reader = csv.reader(f)
    your_list = list(reader)
sampleData = your_list

# Start the demo.
print "===\n"

# Use the IBM Cloudant library to create an IBM Cloudant client.
client = Cloudant(serviceUsername, servicePassword, url=serviceURL)

# Connect to the server
client.connect()

# 2.  Creating a database within the service instance.

# Create an instance of the database.
myDatabaseDemo = client.create_database(databaseName)

# Check that the database now exists.
if myDatabaseDemo.exists():
    print "'{0}' successfully created.\n".format(databaseName)

# Space out the results.
print "----\n"

# 3.  Storing a small collection of data as documents within the database.

# Create documents using the sample data.
# Go through each row in the array
for document in sampleData:
    # Retrieve the fields in each row.
    address = document[0]
    city = document[1]  
    state = document[2] 
    zip = document[3]   
    price = document[4] 
    sqft = document[5]  
    bedrooms = document[6]  
    bathrooms = document[7] 
    days_on_zillow = document[8]    
    sale_type = document[9] 
    url = document[10]  
        
    # Create a JSON document that represents
    # all the data in the row.
    jsonDocument = {
        "address" : address,
        "city" : city,
        "state" : state,
        "zip" : zip,
        "price" : price,
        "sqft" : sqft,
        "bedrooms" : bedrooms,
        "bathrooms" : bathrooms,
        "days_on_zillow" : days_on_zillow,
        "sale_type" : sale_type,
        "url" : url
    }

    # Create a document using the Database API.
    newDocument = myDatabaseDemo.create_document(jsonDocument)

    # Check that the document exists in the database.
    if newDocument.exists():
        print "Document '{0}' successfully created.".format(address)

# Space out the results.
print "----\n"

# 4.  Retrieving a complete list of the documents.

# Simple and minimal retrieval of the first
# document in the database.
result_collection = Result(myDatabaseDemo.all_docs)
print "Retrieved minimal document:\n{0}\n".format(result_collection[0])

# Simple and full retrieval of the first
# document in the database.
result_collection = Result(myDatabaseDemo.all_docs, include_docs=True)
print "Retrieved full document:\n{0}\n".format(result_collection[0])

# Space out the results.
print "----\n"

# Use an IBM Cloudant API endpoint to retrieve
# all the documents in the database,
# including their content.

# Define the end point and parameters
end_point = '{0}/{1}'.format(serviceURL, databaseName + "/_all_docs")
params = {'include_docs': 'true'}

# Issue the request
response = client.r_session.get(end_point, params=params)

# Display the response content
print "{0}\n".format(response.json())

# Space out the results.
print "----\n"

# All done.
# Time to tidy up.

# 5.  Deleting the database.

# Delete the test database.
try :
    client.delete_database(databaseName)
except CloudantException:
    print "There was a problem deleting '{0}'.\n".format(databaseName)
else:
    print "'{0}' successfully deleted.\n".format(databaseName)

# 6.  Closing the connection to the service instance.

# Disconnect from the server
client.disconnect()

# Finish the demo.
print "===\n"

# Say good-bye.
exit()