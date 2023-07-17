[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# jenkins

## Default Metrics

The following metrics are emitted by default. Each of them can be disabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: false
```

### jenkins.job.commit_delta

The calculation difference of the time job was finished minus commit timestamp

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| s | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| job.name | The name of the associated Jenkins job | Any Str |
| job.status | Shows if the job had passed, or failed | Str: ``failed``, ``success``, ``unknown`` |

### jenkins.job.duration

Show the duration of the job

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| s | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| job.name | The name of the associated Jenkins job | Any Str |
| job.status | Shows if the job had passed, or failed | Str: ``failed``, ``success``, ``unknown`` |

### jenkins.jobs.count

Provides a count of the total number of configured jobs

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {Count} | Gauge | Int |