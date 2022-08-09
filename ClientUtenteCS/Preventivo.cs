using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ClientUtenteCS
{
    internal class Preventivo
    {
        public int Id { get; set; }
        public int IdTicket { get; set; }
        public int IdProfessionista { get; set; }
        public string Descrizione { get; set; }
        public string MaterialiRicambi { get; set; }
        public float Costo { get; set; }
        public string DataOra { get; set; }

    }
}
