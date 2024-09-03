# Table: wordpress_core

Core information about a WordPress instance

## Examples

### List name and URL

```sql
select
  name,
  url
from
  wordpress_core
```

### List namespaces

```sql
select 
  jsonb_array_elements_text(namespaces) as namespaces
from
  wordpress_core
```
