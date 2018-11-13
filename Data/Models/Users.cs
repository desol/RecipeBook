using System;
using System.Collections.Generic;

namespace lhj.Data.Models
{
    public partial class Users
    {
        public int Id { get; set; }
        public string Email { get; set; }
        public string Password { get; set; }
        public short Attempts { get; set; }
        public bool? Locked { get; set; }
        public short Status { get; set; }
        public DateTime Created { get; set; }
        public DateTime? Lastlogin { get; set; }
    }
}
