// Data Models of Inventory Itemization
#include <ctime>
// Basic Setting File Configuration
struct Settings {
    char strHostName[128]; // The host name to call home to
    char strGroupID[64]; // The group ID to sent with the data
    int intFrequency; // How many minutes between sending data
};

struct InventoryV1 {
    std::time_t objTimestamp; // Colect the timestamp of when this information was collected
    double dblLatitude; // Latitude of this device
    double dblLongitude; // Longitude of this device
    
    char strMake[128]; // Make of device - Samsung, Dell, HP, etc. - who made this thing
    char strModel[128]; // Model of the device - HP240Z, PK202, etc. - what is this specific device
    char strSerial[128]; // Serial number of the device - what is unique about this device in phone cases this is the IMEI
    char strOSVersion[32]; // What version of the operating system is running on this device
    char strOSStatus[64]; // What is the status of this device OS, phones are is it rooted or not
    char strPhysicalCondition[64]; // What if anything is physically wrong with the device if that information is available
};