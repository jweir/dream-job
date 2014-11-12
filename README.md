# Dream Job 
A job queue and scheduler.

Not Resque, but influenced by it

What is wrong with Resque & Resque-Scheduler, why not use them?

# Major Goals
A crashed scheduler will know when it last ran.  It will make up for lost time
and rerun skipped jobs, if the job is speced to do so.

Workers will be sent a quit signal, and restarted only after they have finished their current work.

Having something like the Resque failure backends.

## Secondary Goals
Mechanism for job tracking
  - how long did a job take
  - jobs success / failure

Jobs get queued
- at run time Enqueue()
- scheduled EnqueueAt(Time)
- dependent EnqueueAfter(JobHandle)

Jobs should be designed so they are atomic.
  For example a job which downloads, writes and imports into a DB would
  have 3 discrete steps.  A failure at one would not require rerunning the 
  previous steps.

Visualization to see what jobs ran and what jobs are coming up.


### The Job handle

This is the unique id for each job.  Jobs will have a name, args and the time scheduled to run (if not then the time they were enqueued).

   [ name, time, args(used to initiate the job) ]

EnqueueAfter will require some additional logic for the job handle.

  EnqueueAfter(JobHandle(GetPrices, 16:00))

### Job Args

The arguments used to initiate a job should be kept very simple.  Passing in complex structures is a smell.
All arguments should be JSON compatible.


