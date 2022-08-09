using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ClientUtenteCS
{
    internal class Ticket
    {
        public int Id { get; set; }
        public string Stato { get; set; }
        public string Categoria { get; set; }
        public string Titolo { get; set; }
        public string Descrizione { get; set; }

        public override string ToString()
        {
            return "Id: " + Id + "stato: " + Stato + "categoria: " + Categoria + "titolo: " + Titolo;
        }
    }
}
