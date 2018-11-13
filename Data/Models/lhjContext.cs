using System;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata;

namespace lhj.Data.Models
{
  public partial class lhjContext : DbContext
  {
    public lhjContext()
    {
    }

    public lhjContext(DbContextOptions<lhjContext> options)
        : base(options)
    {
    }

    public virtual DbSet<Harmontown> Harmontown { get; set; }
    public virtual DbSet<Users> Users { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
    {
    }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
      modelBuilder.Entity<Harmontown>(entity =>
      {
        entity.ToTable("harmontown");

        entity.HasIndex(e => e.Title)
                  .HasName("harmontown_un")
                  .IsUnique();

        entity.Property(e => e.Id).HasColumnName("id");

        entity.Property(e => e.Date)
                  .HasColumnName("date")
                  .HasColumnType("timestamp with time zone");

        entity.Property(e => e.Listens).HasColumnName("listens");

        entity.Property(e => e.Location)
                  .IsRequired()
                  .HasColumnName("location");

        entity.Property(e => e.Title)
                  .IsRequired()
                  .HasColumnName("title");
      });

      modelBuilder.Entity<Users>(entity =>
      {
        entity.ToTable("users");

        entity.Property(e => e.Id).HasColumnName("id");

        entity.Property(e => e.Attempts).HasColumnName("attempts");

        entity.Property(e => e.Created)
                  .HasColumnName("created")
                  .HasColumnType("timestamp with time zone")
                  .HasDefaultValueSql("now()");

        entity.Property(e => e.Email)
                  .IsRequired()
                  .HasColumnName("email");

        entity.Property(e => e.Lastlogin)
                  .HasColumnName("lastlogin")
                  .HasColumnType("timestamp with time zone");

        entity.Property(e => e.Locked)
                  .IsRequired()
                  .HasColumnName("locked")
                  .HasDefaultValueSql("true");

        entity.Property(e => e.Password)
                  .IsRequired()
                  .HasColumnName("password");

        entity.Property(e => e.Status).HasColumnName("status");
      });
    }
  }
}
