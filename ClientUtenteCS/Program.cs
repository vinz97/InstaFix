using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Diagnostics;
using System.Text;
using System.Net.Http;

namespace ClientUtenteCS
{
    internal static class Program
    {
        /// <summary>
        /// Punto di ingresso principale dell'applicazione (main)
        /// </summary>
        [STAThread]
        static void Main()
        {
         Application.EnableVisualStyles();
         Application.SetCompatibleTextRenderingDefault(false);
         Application.Run(new FormLogin());
        }


        /*
         *     //    Richieste.HttpGet("http://localhost:8000/hello");
                 Richieste.HttpGetAsync("http://localhost:8000/hello").Wait();

                    string json = "hello";
                    var stringContent = new StringContent(json, Encoding.UTF8, "application/json");
                    Richieste.HttpPostSync("http://localhost:8000/upload", stringContent).Wait();

         */
    }
}
