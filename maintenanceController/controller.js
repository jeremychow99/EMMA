import express from 'express'
import cors from 'cors'
import request from 'request'


const app = express();
const port = 8080;

app.use(cors());
app.use(express.json())

//  Schedule Maintenance
app.post('/schedule_maintenance', (req, res) => {
    console.log(req.body)
    // 1. Reserve Parts (RabbitMQ)

    // 2. Create Maintenance Record
    

    // 3. Send out Notification

    res.json({ msg: "Successfully scheduled maintenance" });
})


// Start Maintenance
app.patch('/start_maintenance/:maintenanceId', (req, res) => {
    console.log(req.body)
    // 1. Send update maintenance (Start time, TechId)
    // 2. Send out Notification


    res.json({ msg: "Successfully started maintenance" });
})


// Request for More Parts
app.patch('/request_parts/:maintenanceId', (req, res) => {
    console.log(req.body)
    // 1. Reserve Parts (HTTP)
    // 2. Update PartList in Maintenance (Successfully reserved parts)


    res.json({ msg: "Parts has been requested" });
})


// End Maintenance
app.patch('/end_maintenance/:maintenanceId', (req, res) => {
    console.log(req.body)
    // 1. Send update maintenance (Partlist, EndTime, Desc)
    // 2. Return Parts (RabbitMQ)
    // 3. Send out Notification 


    res.json({ msg: "Successfully ended maintenance" });
})


app.listen(port, () => {
    console.log(`Maintenance Controller listening on port ${port}`)
})


