using System;
using System.Collections.Generic;

namespace lhj.Data.Models
{
    public partial class Harmontown
    {
        public int Id { get; set; }
        public string Title { get; set; }
        public DateTime Date { get; set; }
        public string Location { get; set; }
        public short Listens { get; set; }
    }
}
