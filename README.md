# hasher2 package

  

How to use:

  

1. Import package

  

```

import (

"github.com/d-tanasienko/hasher"

)

```

  

2. Create a hash from a password:

  

```

hasher.HashPassword('password')

```

  

3. Check that hash matches the input password:

  

```

hasher.CheckPasswordHash('password', 'hash')

```
