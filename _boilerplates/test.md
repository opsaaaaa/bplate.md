---
layout: post

    
_boilerplate:

  timestamp: false

  type: testing
  title: title
  num: 130 

  path: test
     
author: John Doe
title: {{ boilerplate.title }}

created: {{ boilerplate.time }}
random_url: {{ boilerplate.random }}
random_url: {{ boilerplate.random=5 }}
custom: {{ boilerplate.custom }}
num: {{ boilerplate.num }}
---


Default Test Heading
--------------------

some boilerplate text

Type is {{ boilerplate.type }}
File is {{ boilerplate.file }}
