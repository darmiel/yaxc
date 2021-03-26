package server

const body = `
 __  __     ______     __  __     ______    
/\ \_\ \   /\  __ \   /\_\_\_\   /\  ___\   
\ \____ \  \ \  __ \  \/_/\_\/_  \ \ \____  
 \/\_____\  \ \_\ \_\   /\_\/\_\  \ \_____\ 
  \/_____/   \/_/\/_/   \/_/\/_/   \/_____/ 
                                            
Just POST your contents to /:anywhere
POST /:anywhere your contents
GET /:anywhere for your contents

GET /hash/:anywhere for content hash
POST /:anywhere?ttl=3m for custom TTL

POST /:anywhere?secret=password to protect your contents
GET /:anywhere?secret=password to get protected contents`
