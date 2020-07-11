---
title: "Authenticate"
date: 2020-07-10T21:21:05-03:00
draft: false
---

<div data-netlify-identity-menu></div>

<script type="text/javascript" src="https://identity.netlify.com/v1/netlify-identity-widget.js"></script>

<script type="text/javascript">
const user = netlifyIdentity.currentUser();
if (user) console.log(user);
</script>