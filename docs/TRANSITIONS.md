# Transitions docs

all transitions require the following fields:

```json5
{
   
  "trackingId": "", // <STRING><REQUIRED> 99minutos guide number
  "eventName":  "", // <STRING><REQUIRED> event name in upper case

  // metadata for internal use
  "metadata": {
    "platform": "", // <STRING><REQUIRED> name of the platform that generates the event like: 
                    // api-v3|fulfill-gateway|induction-bff|stations-service
    ...
  },
  
  // specific data required for another's systems in 99minutos
  "data": {                 
    "commentClient":   "", // <STRING><OPTIONAL> comment from the client { VISIBLE IN THE TRACKING PAGE }
    "commentInternal": "", // <STRING><OPTIONAL> comment from 99minutos { VISIBLE IN CONSOLE }
  },
  
  // location of the event
  "location": {
    "id": "",         // <STRING> UUId from geo-location
    "type": "",       // <STRING> private|p99
    "category": "",   // <STRING> dropoff|pickup
    "latitude": 0.0,  // <FLOAT>  between -90 and 90
    "longitude": 0.0  // <FLOAT>  between -180 and 180
  },
  
  // actor of the event
  "actor": {
    "name": "",    // <STRING> examples:
    "type": ""     // <STRING> type: operator  actor:  mx2 someone@99minutos.com
                   //          type: system    actor:  schedule|task|manual
                   //          type: driver    actor:  ios|android|web <version>
                   //          type: customer  actor:  api|integration|web|mobile
                   //          type: unknown   actor:  free message for identification
  },
  "eventTime": ""  // <STRING><REQUIRED><UTC><RFC3339> always use UTC '2023-01-01T00:00:00.00000Z'  NOTE las 5 decimales NO 3
}
```

# Data specs

## Creation

---

### INIT ▶ 1001 - NEW_DRAFT

Emited by api v3

```json5
{
    "data": {  } // NOT DEFINED
}
```

### INIT ▶ 1002 - NEW_ORDER_CONFIRMED

Emited by api v3

```json5
{
    "data": {  } // NOT DEFINED
}
```

### 1001 ▶ 1002 - DRAFT_CONFIRMED

```json5
{
    "data": {  } // NOT DEFINED
}
```

### 1002 ▶ 2001 - TO_PICKUP

```json5
{
    "data": {
        "staff_id": "", // <STRING> staff id from 99minutos 
    }
}
```

### 1002 ▶ 3001 - DIRECT_TO_STORE

```json5
{
    "data": {
        "station": "", // <STRING> MX1|MX2|TPC1|GDL1|ETC...
    }
}
```

## Pickup

---

### 2001 ▶ 2002 - DRIVER_ASSIGNED_TO_PICKUP

```json5
{
    "data": {
        "staff_id": "",            // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "vehicle_id": "",          // <STRING> vehicle id from 99minutos
        "session_id": "",          // <STRING> session id from 99minutos
        "pickup": {            
            "routing_id": "",      // <STRING> routing id from 99minutos
            "idempotence_key": "", // <STRING> idempotence key from 99minutos
            "group": "",           // <STRING> group id ( name of the group/prefix )
            "point": 0,            // <INT> point id
            "secuence": 0,         // <INT> secuence id
        },
    }
}
```

### 2002 ▶ 2001 - DRIVER_CANCELLATION

```json5
{
    "data": {
        "commentInternal": "", // <STRING> reason of cancellation
    }
}
```

### 2002 ▶ 2002 - DRIVER_REASSIGNED_TO_PICKUP

```json5
{
    "data": {
        "commentInternal": "", // <STRING> reason of reassignment
        "staff_id": "",        // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "vehicle_id": "",      // <STRING> vehicle id from 99minutos
    }
}
```

### 2002 ▶ 2003 - PICKUP_CONFIRMED

```json5
{
    "data": {
        "staff_id": "",       // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "pickup_order": 0,    // <INT> position in which it was collected
        "location": {
            "latitude": 0.0,  // <FLOAT> between -90 and 90
            "longitude": 0.0  // <FLOAT> between -180 and 180
        },
    }
}
```

### 2002 ▶ 2101 - PICKUP_FAILED

```json5
{
    "data": {
        "commentInternal": "", // <STRING> reason of reassignment
        "staff_id": "",        // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "location": {
            "latitude": 0.0,   // <FLOAT> between -90 and 90
            "longitude": 0.0   // <FLOAT> between -180 and 180
        },
    }
}
```

### 2101 ▶ 2001 - NEW_PICKUP_ATTEMPT

```json5
{
    "data": {
        "staff_id": "",     // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "vehicle_id": "",   // <STRING> vehicle id from 99minutos
    }
}
```

### 2001 ▶ 8201 - MAX_ATTEMPTS_EXCEDED

```json5
{
    "data": {  } // NOT DEFINED
}
```

### 2003 ▶ 3001 - TO_STORE

```json5
{
    "data": { 
        "station": "",  // <STRING> MX1|MX2|TPC1|GDL1|ETC...
     }
}
```

## Warehouse & Linehaul

---

### 3001 ▶ 3002 - CONTAINER_ASSIGNED

```json5
{
    "data": {
        "delivery": {              // only when container_type is LASTMILE  
            "routing_id": "",      // <STRING> routing id from 99minutos
            "idempotence_key": "", // <STRING> idempotence key from 99minutos
            "group": "",           // <STRING> group id
            "point": 0,            // <INT> point id
            "secuence": 0,         // <INT> secuence id
        },
        "station": "",             // <STRING> MX1|MX2|TPC1|GDL1|ETC...
        "container_id": "",        // <STRING> container id
        "container_tag": "",       // <STRING> container tag
        "container_type": ""       // <STRING> MIDDLE|LASTMILE 
    }
}
```

### 3002 ▶ 3002 - CONTAINER_REASSIGNED

```json5
{
    "data": { 
        "commentInternal": "", // <STRING> reason of reassignment
        "staff_id": "",        // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "vehicle_id": "",      // <STRING> vehicle id from 99minutos
        "station": "",         // <STRING> MX1|MX2|TPC1|GDL1|ETC...
    }
}
```

### 3002 ▶ 3001 - CONTAINER_CANCELLED

```json5
{
    "data": { 
        "commentInternal": "", // <STRING> reason of cancellation
        "station": "",         // <STRING> MX1|MX2|TPC1|GDL1|ETC...
    }
}
```

### 3002 ▶ 3003 - VEHICLE_ASSIGNED [ DEPRECATED ]

```json5
{
    "data": {
        "vehicle_id": "",    // <STRING> vehicle id from 99minutos
        "container_id": "",  // <STRING> container id
        "container_tag": "", // <STRING> container tag
    }
}
```

### 3003 ▶ 3004 - TRANSFER_REQUESTED

```json5
{
    "data": {
        "station": "", // <STRING> MX1|MX2|TPC1|GDL1|ETC...
    }
}
```

### 3004 ▶ 3001 - NEW_STORAGE_POINT

```json5
{
    "data": {
        "station": "", // <STRING>        
    }
}
```

### 3001 ▶ 7401 - MAX_ATTEMPTS_EXCEDED_DELIVERY

The process is considered automatic

### 3001 ▶ 7501 - MAX_ATTEMPTS_EXCEDED_RETURN

The process is considered automatic

### 3002 ▶ 4001 - CONTAINER_ASSIGNED_TO_DELIVERY

```json5
{
    "data": {
        "staff_id": "",      // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "vehicle_id": "",    // <STRING> vehicle id from 99minutos
        "container_id": "",  // <STRING> container id
        "container_tag": "", // <STRING> container tag
    }
}
```

### 3002 ▶ 5001 - CONTAINER_ASSIGNED_TO_RETURN

```json5
{
    "data": {
        "staff_id": "",      // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "vehicle_id": "",    // <STRING> vehicle id from 99minutos
        "container_id": "",  // <STRING> container id
        "container_tag": "", // <STRING> container tag
    }
}
```

### 3003 ▶ 4001 - DRIVER_ASSIGNED_TO_DELIVERY [ DEPRECATED ]

```json5
{
    "data": {  
        "staff_id": "",  // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
    }
}
```

### 3003 ▶ 5001 - DRIVER_ASSIGNED_TO_RETURN [ DEPRECATED ]

```json5
{
    "data": {  
        "staff_id": "", // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
    }
}
```

### 3001 ▶ 3001 - SWITCH_IN_STORE

```json5
{
    "data": {
        "commentInternal": "", // <STRING> reason of return
        "isReturn": false      // <BOOLEAN> true if is return 
    }
}
```

## Delivery

---

### 4001 ▶ 4001 - DRIVER_REASSIGNED_TO_DELIVERY

```json5
{
    "data": { 
        "staff_id": "",       // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "commentInternal": "",// <STRING> reason of reassignment
    }
}
```

### 4001 ▶ 4002 - DELIVERY_CONFIRMED

```json5
{
    "data": {
        "evidence": "",       // <STRING> URL to evidence 
        "location": {
            "latitude": 0.0,  // between -90 and 90
            "longitude": 0.0  // between -180 and 180
        },
    }
}
```

### 4001 ▶ 4003 -

```json5
{
    "data": {  } // NOT DEFINED
}
```

### 4003 ▶ 4002 -

```json5
{
    "data": {  } // NOT DEFINED
}
```

### 4001 ▶ 4101 - DELIVERY_FAILED

```json5
{
    "data": {
        "commentInternal": "", // <STRING> reason of failure
        "evidence": "",        // <STRING> URL to evidence 
        "location": {
            "latitude": 0.0,   // between -90 and 90
            "longitude": 0.0   // between -180 and 180
        },
    }
}
```

### 4101 ▶ 3001 - NEW_DELIVERY_ATTEMPT

```json5
{
    "data": {
        "station": "",  // <STRING> MX1|MX2|TPC1|GDL1|ETC...
    }
}
```

## Return

---

### 5001 ▶ 5001 - DRIVER_REASSIGNED_TO_RETURN

```json5
{
    "data": {
        "staff_id": "",        // <STRING> driver staff id from 99minutos
        "staff": {
            "driver_id": "",       // id from legacy 
            "staff_id": "",        // id from staff
            "name": "",
            "nickname": "", 
            "type": "",            // base | freelancer | from_provider
            "provider": {          // only when type is from_provider
                "staff_id": "",
                "name": "",
                "email": "",
                "trade_name": "",
                "nomenclature": "",
            }, 
        },
        "commentInternal": "", // <STRING> reason of reassignment
    }
}
```

### 5001 ▶ 5002 - RETURN_CONFIRMED

```json5
{
    "data": {
        "evidence": "",       // <STRING> URL to evidence 
        "location": {
            "latitude": 0.0,  // between -90 and 90
            "longitude": 0.0  // between -180 and 180
        },
    }
}
```

### 5001 ▶ 5101 - RETURN_FAILED

```json5
{
    "data": {
        "commentInternal": "", // <STRING> reason of failure
        "evidence": "",        // <STRING> URL to evidence 
        "location": {
            "latitude": 0.0,   // between -90 and 90
            "longitude": 0.0   // between -180 and 180
        },
    }
}
```

### 5101 ▶ 3001 - NEW_RETURN_ATTEMPT

```json5
{
    "data": {
        "station": "", // <STRING> MX1|MX2|TPC1|GDL1|ETC...
    }
}
```




