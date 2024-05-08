# Routes
For seeing post body parameters visit postman.This file is only explaining what each route does.
## Agents

    GET /api/agents/{ID} => Get One by ID
    POST /api/agents/ => Create agent 

## Orders
    GET /api/agents/{ID} => Get One by ID
    POST /api/agents/ => Create 

## Vendors
    GET /api/vendors/{ID} => Get One by ID
    POST /api/vendors/ => Create
## Trips
    GET /api/trips/by-order-id/{ID} => Get One by Order ID
    POST /api/trips/ => Create
    POST /api/trips/change-status => Update Status of trip

## Support
    GET /api/support/order-status/{ID} => See DealyReports for Order (can be used by end user)
    POST /api/support/report-order-delay => Customer can use this route to report delay on their order.

    GET /api/support/get-task-for-agent/{AgentID} => Agent can request for a task using this.If already has a task assigned to them will get the already assigned task instead of a new one.
    POST /api/support/update-task-for-agent/ => Agent can update their task status here.
## Analytics
    GET /api/analytics/vendor-delays-weekly-report => Vendor total delays this week.

## The task requested flow:
- User can create order using POST /orders
- A trip for that order can be made using POST /trips
- User can report delay using /support/report-order-delay
- Agent can check for tasks using /get-task-for-agent
